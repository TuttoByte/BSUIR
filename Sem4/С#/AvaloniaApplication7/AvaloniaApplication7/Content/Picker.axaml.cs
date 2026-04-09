using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text.RegularExpressions;
using Avalonia;
using Avalonia.Controls;
using Avalonia.Controls.Primitives;
using Avalonia.Markup.Xaml;
using AvaloniaApplication7.Services.Contracts;
using AvaloniaApplication7.Services.Entities;
using AvaloniaApplication7.Services.Models;

namespace AvaloniaApplication7.Content;

public partial class Picker : UserControl, INotifyPropertyChanged
{


    public ObservableCollection<Singer> Groups { get; set; } = new();
    public ObservableCollection<Song> Songs { get; set; } = new();
    
    
    private Singer? _selectedSinger;

    public Singer? SelectedSinger
    {
        get => _selectedSinger;
        set
        {
            if (SetProperty(ref _selectedSinger, value))
            {
                if (value != null)
                {
                    LoadFromGroup(value.Id);
                }
            }
           
        }
    }
    
    public IDbService _db;
    public Picker()
    {
        InitializeComponent();
        DataContext = this;
        _db = new SqLiteService();
        _db.Init();
        LoadGroup();
        
    }
    
    
    public string SongCountText
    {
        get
        {
            return  $"Количество песен: {Songs.Count}";
        }
    }

    public void LoadFromGroup(int  groupId)
    {
       var songs = _db.GetSingerSongs(groupId);
       Songs.Clear();
       foreach (var song in songs)
       {
           Songs.Add(song);
       }

     
    }
    
    public void LoadGroup()
    {
        var singers = _db.GetAllSingers();
        foreach (var singer in singers)
        {
            Groups.Add(singer);
        }
    }
    
    public event PropertyChangedEventHandler? PropertyChanged;

    protected bool SetProperty<T>(ref T field, T value, [CallerMemberName] string? propertyName = null)
    {
        if (EqualityComparer<T>.Default.Equals(field, value)) return false;
        
        field = value;
        PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
        return true;
    }

    protected void OnPropertyChanged([CallerMemberName] string propertyName = null)
    {
        PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
    }
}


